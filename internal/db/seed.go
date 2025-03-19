package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"math/rand"

	"github.com/wadiya/go-social/internal/store"
)

var userNames = []string{
	"Aditya", "Janardhan", "Akshay", "Prachi", "Neha", "Nikhil", "Pratik", "Saurabh", "Rohan", "Tanvi",
	"Vikas", "Swati", "Kunal", "Pooja", "Aniket", "Sneha", "Rahul", "Megha", "Amit", "Komal",
	"Vivek", "Ishita", "Siddharth", "Gauri", "Rajat", "Rupali", "Tushar", "Shweta", "Yash", "Deepika",
	"Harsh", "Reshma", "Sandeep", "Manisha", "Chetan", "Rina", "Omkar", "Pallavi", "Varun", "Mitali",
	"Shivam", "Anushka", "Rakesh", "Kiran", "Bhavna", "Suhas", "Ruchi", "Dinesh", "Asmita", "Ganesh",
}

var titles = []string{
	"Greatest Cricket Matches Ever",
	"Top 10 Batsmen of All Time",
	"How to Play a Perfect Cover Drive",
	"Legendary Cricket Captains",
	"Fastest Centuries in Cricket History",
	"Spin Bowling Masterclass",
	"Understanding DLS Method in Cricket",
	"Top 5 IPL Moments of All Time",
	"How to Improve Your Batting Skills",
	"Greatest Rivalries in Cricket",
	"Famous Cricket World Cup Upsets",
	"Best All-Rounders in Cricket History",
	"Secrets of Reverse Swing Bowling",
	"Evolution of T20 Cricket",
	"Best Wicket-Keepers in Cricket",
	"Why Test Cricket is Still King",
	"Iconic Cricket Commentary Moments",
	"Most Dramatic Last Over Finishes",
	"Top 10 Fielding Moments in Cricket",
	"The Rise of Womens Cricket",
	"Memorable Cricket World Cup Finals",
	"Cricket Strategies for Winning T20s",
	"Best Yorker Bowlers of All Time",
	"Top 5 Cricketing Controversies",
	"Captaincy Tactics in ODI Cricket",
	"History of the Ashes Series",
	"How to Bowl the Perfect Bouncer",
	"Most Shocking Cricket Injuries",
	"Famous Cricket Comebacks",
	"Best Cricket Stadiums in the World",
}

var contents = []string{
	"A look at the most thrilling and unforgettable cricket matches that kept fans on the edge of their seats.",
	"A ranking of the greatest batsmen in cricket history, based on records, technique, and impact on the game.",
	"Learn the fundamentals of executing a flawless cover drive like cricketing legends Virat Kohli and Sachin Tendulkar.",
	"A tribute to the most inspiring cricket captains who led their teams to glory with skill and strategy.",
	"Exploring the top players who smashed the fastest centuries in Tests, ODIs, and T20s.",
	"A breakdown of spin bowling techniques, from leg-spin to off-spin, used by legendary bowlers.",
	"A simplified explanation of the Duckworth-Lewis-Stern method used in rain-affected matches.",
	"A countdown of the most electrifying and unforgettable moments in the Indian Premier League.",
	"Essential tips to enhance batting technique, footwork, and shot selection.",
	"The most intense cricket rivalries that have defined the sport over decades.",
	"Unexpected results that shocked the world in the history of Cricket World Cups.",
	"A tribute to players who excelled both with the bat and ball in international cricket.",
	"How pacers use reverse swing to outsmart batsmen in challenging conditions.",
	"How the shortest format of cricket has revolutionized the game and attracted new fans.",
	"A look at the finest wicket-keepers known for their agility, sharp reflexes, and leadership.",
	"Despite the rise of T20s, Test cricket remains the ultimate format for purists and players.",
	"The most memorable and emotional moments in cricket history, narrated by legendary commentators.",
	"Heart-stopping matches that were decided in the final over, leaving fans breathless.",
	"Stunning catches, brilliant run-outs, and extraordinary fielding efforts that changed games.",
	"How womens cricket has grown in popularity and recognition worldwide.",
	"A review of the most thrilling World Cup finals that shaped the sports history.",
	"Key strategies that teams use to dominate in T20 cricket and secure victories.",
	"Celebrating the greatest bowlers who mastered the deadly yorker delivery.",
	"The most controversial moments that sparked debates and shook the cricketing world.",
	"How captains devise strategies, set fields, and lead their teams to success in ODIs.",
	"The legendary England-Australia rivalry that has defined Test cricket for over a century.",
	"A guide to bowling lethal bouncers like Mitchell Starc and Jofra Archer.",
	"Unfortunate injuries that stunned fans and changed the careers of top cricketers.",
	"Incredible comebacks by teams and players who defied odds to achieve greatness.",
	"A tour of the most iconic and picturesque cricket stadiums across the globe.",
}

var tags = []string{
	"Cricket", "Great Matches", "Thrillers", "Historic Games",
	"Batsmen", "Cricket Legends", "Top Players", "Records",
	"Batting Tips", "Cover Drive", "Cricket Techniques", "Shot Making",
	"Captains", "Leadership", "Cricket Strategy", "Inspirational Players",
	"Fastest Centuries", "Records", "Batting Feats", "Cricket History",
	"Spin Bowling", "Bowling Techniques", "Cricket Skills", "Leg Spin",
	"DLS Method", "Rain Rules", "Cricket Calculations", "Match Impact",
	"IPL", "T20 Cricket", "Best Moments", "Cricket Leagues",
	"Batting Improvement", "Cricket Training", "Shot Selection", "Footwork",
	"Rivalries", "Cricket Feuds", "Greatest Matches", "Historic Battles",
	"Cricket World Cup", "Upsets", "Surprise Wins", "Memorable Matches",
	"All-Rounders", "Cricket Stars", "Batting & Bowling", "Match Winners",
	"Reverse Swing", "Bowling Mastery", "Pace Attack", "Swing Bowling",
	"T20 Evolution", "Modern Cricket", "Cricket Growth", "Game Changing",
	"Wicket-Keeping", "Best Keepers", "Cricket Reflexes", "Legendary Players",
	"Test Cricket", "Purest Form", "Long Format", "Ultimate Challenge",
	"Cricket Commentary", "Iconic Calls", "Best Moments", "Legendary Voices",
	"Last Over Finishes", "Close Matches", "Thrilling Games", "Final Over Drama",
	"Fielding Masterclass", "Best Catches", "Run Outs", "Agility",
	"Womens Cricket", "Rise of Womens Game", "Cricket Growth", "Top Players",
	"World Cup Finals", "Cricket History", "Big Matches", "Unforgettable Games",
	"T20 Strategies", "Winning Tactics", "Cricket Strategy", "Game Plans",
	"Yorker Specialists", "Best Bowlers", "Pace Bowling", "Cricket Domination",
	"Cricket Controversies", "Scandals", "Shocking Moments", "Unforgettable Incidents",
	"Captaincy", "Cricket Leadership", "Game Planning", "Tactical Moves",
	"Ashes", "Historic Rivalry", "England vs Australia", "Test Cricket",
	"Bouncers", "Pace Bowling", "Aggressive Bowling", "Bowling Tips",
	"Cricket Injuries", "Shocking Moments", "Career Impact", "Painful Incidents",
	"Comebacks", "Greatest Returns", "Cricket Redemption", "Fighting Spirit",
	"Cricket Stadiums", "Best Grounds", "Historic Venues", "Beautiful Stadiums",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return

		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num) // make() is a built-in Go funtion

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: userNames[i%len(userNames)] + fmt.Sprintf("%d", i),
			Email:    userNames[i%len(userNames)] + fmt.Sprintf("%d", i) + "@example.com",
			// Password: "123456",
		}
	}

	return users

}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}
