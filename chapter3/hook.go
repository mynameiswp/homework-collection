package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 模型定义
type User struct {
	gorm.Model
	Name         string `gorm:"size:255;not null"`
	ArticleCount int    `gorm:"default:0;comment:用户发布的文章数量"`
	Posts        []Post // 一对多关联：用户-文章
}

type Post struct {
	gorm.Model
	Title         string    `gorm:"size:255;not null"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"index;comment:关联用户ID"`
	CommentStatus string    `gorm:"default:'有评论';size:20;comment:评论状态（有评论/无评论）"`
	Comments      []Comment // 一对多关联：文章-评论
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"index;comment:关联文章ID"`
}

// 2. 钩子函数：创建文章时更新用户文章数
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("article_count", gorm.Expr("article_count + 1")).Error; err != nil {
		return fmt.Errorf("更新用户文章数量失败: %v", err)
	}
	fmt.Printf("[钩子函数] 文章《%s》创建，用户ID:%d 文章数+1\n", p.Title, p.UserID)
	return nil
}

// 3. 钩子函数：删除评论后检查并更新文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 步骤1：获取当前评论所属文章
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return fmt.Errorf("获取文章失败: %v", err)
	}

	// 步骤2：统计该文章剩余评论数
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return fmt.Errorf("统计评论数量失败: %v", err)
	}

	// 步骤3：根据评论数更新状态
	oldStatus := post.CommentStatus
	if commentCount == 0 {
		post.CommentStatus = "无评论"
	} else {
		post.CommentStatus = "有评论" 
	}

	if err := tx.Save(&post).Error; err != nil {
		return fmt.Errorf("更新文章评论状态失败: %v", err)
	}

	// 打印钩子函数执行日志（动态过程）
	fmt.Printf("[钩子函数] 评论ID:%d 删除 → 文章《%s》剩余评论数:%d → 状态从「%s」变为「%s」\n",
		c.ID, post.Title, commentCount, oldStatus, post.CommentStatus)
	return nil
}

// 输出文章详情
func printPostDetail(db *gorm.DB, postID uint) {
	var post Post
	db.Preload("Comments").First(&post, postID)
	fmt.Printf("\n【文章详情】ID:%d 标题《%s》\n", post.ID, post.Title)
	fmt.Printf("  评论状态：%s | 当前评论数：%d\n", post.CommentStatus, len(post.Comments))
	if len(post.Comments) > 0 {
		fmt.Print("  评论列表：")
		for i, c := range post.Comments {
			if i > 0 {
				fmt.Print(" | ")
			}
			fmt.Printf("评论%d:%s", c.ID, c.Content)
		}
		fmt.Println()
	} else {
		fmt.Println("  评论列表：无")
	}
	fmt.Println("----------------------------------------")
}

//输出用户-文章-评论关系
func printUserPostsWithComments(user User) {
	fmt.Printf("\n===== 用户「%s」(ID:%d) 的所有文章 =====\n", user.Name, user.ID)
	for _, post := range user.Posts {
		fmt.Printf("文章ID:%d 标题《%s》 | 评论状态：%s | 评论数：%d\n",
			post.ID, post.Title, post.CommentStatus, len(post.Comments))
	}
	fmt.Println("========================================")
}

func main() {
	// 连接MySQL数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接MySQL失败: " + err.Error())
	}

	// 配置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移表结构
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	

	//清空旧数据
	db.Where("1=1").Delete(&Comment{})
	db.Where("1=1").Delete(&Post{})
	db.Where("1=1").Delete(&User{})
	fmt.Println("✅ 旧测试数据已清空")

	
	// 造数
	users := []User{{Name: "张三"}, {Name: "李四"}}
	db.Create(&users)
	
	
	posts := []Post{
		{Title: "张三的Go入门教程", Content: "Go的基础语法讲解", UserID: users[0].ID},
		{Title: "张三的MySQL优化", Content: "索引优化实战", UserID: users[0].ID},
		{Title: "李四的Python爬虫", Content: "Scrapy框架使用", UserID: users[1].ID},
		{Title: "李四的前端笔记", Content: "Vue3+TS实战", UserID: users[1].ID},
	}
	db.Create(&posts)
	

	comments := []Comment{
		// 张三的文章1：3条评论（重点测试删除）
		{Content: "讲得太清楚了！", PostID: posts[0].ID},
		{Content: "请问Go的协程怎么理解？", PostID: posts[0].ID},
		{Content: "收藏了，谢谢博主", PostID: posts[0].ID},
		// 张三的文章2：1条评论
		{Content: "索引优化太实用了", PostID: posts[1].ID},
		// 李四的文章1：0条评论
		// 李四的文章2：2条评论
		{Content: "Vue3的组合式APIyyds", PostID: posts[3].ID},
		{Content: "TS类型校验太香了", PostID: posts[3].ID},
	}
	db.Create(&comments)

	
	// 查询张三的所有文章及其对应的评论信息
	var userZhang User
	db.Preload("Posts.Comments").First(&userZhang, users[0].ID)
	printUserPostsWithComments(userZhang)
	
	printPostDetail(db, posts[0].ID)

	// ========== 步骤4：分步删除评论，测试钩子函数动态过程 ==========
	fmt.Println("\n========== 开始删除评论，测试钩子函数 ==========")

	// 4.1 删除张三文章1的第1条评论（剩余2条，状态仍为“有评论”）
	
	db.Delete(&comments[0])
	printPostDetail(db, posts[0].ID)

	// 4.2 删除张三文章1的第2条评论（剩余1条，状态仍为“有评论”）

	db.Delete(&comments[1])
	printPostDetail(db, posts[0].ID)

	// 4.3 删除张三文章1的最后1条评论（剩余0条，触发状态变为“无评论”）
	
	db.Delete(&comments[2])
	printPostDetail(db, posts[0].ID)

	// ========== 步骤5：最终状态验证 ==========
	
	// 重新查询张三的所有文章，确认状态变化
	db.Preload("Posts.Comments").First(&userZhang, users[0].ID)
	printUserPostsWithComments(userZhang)

}
