# 💸 Budget Tracker – Personal Project

This is a small side project I'm building to help me track my daily expenses and manage my budget more clearly.  
I'm creating it using **Golang (Gin)** for the backend and **PostgreSQL** for the database.  
Nothing too fancy — just something simple, useful, and something I can keep improving as I learn more.

Right now the project is still in progress, and I'm building it step by step whenever I have free time.

---

## 🎯 Why I Built This

Mostly because:
- I want a tool that helps me understand where my money goes  
- I enjoy building things that can actually be used in real life  
- I want to improve my backend skills, especially with Go, structuring APIs, and working with databases  

Plus… I just love the feeling of making my own tools instead of relying on apps I don’t fully control 😄

---

## ⚙️ Tech Stack

- **Golang (Gin Framework)** — for building the API  
- **PostgreSQL** — to store transactions, categories, and budgets  
- **Gorm** — ORM for database interactions  
- **JSON API** — for future integration with a frontend (maybe Astro or Next.js later)

---

## 📦 Features (Current & Upcoming)

### ✅ Already Working
- Add budget records  
- Add daily expenses  
- Store everything in PostgreSQL  
- Basic CRUD API  
- Simple category management  

### 🚧 In Progress
- Monthly budget summary  
- Auto-calculation of remaining balance  
- Sorting & filtering  
- Better structuring for services & repository layers  

### 💡 Future Ideas
- Web dashboard  
- Export to CSV or PDF  
- Notifications (if I feel ambitious 😄)

---

## 🛠️ How to Run the Project

```bash
# Clone the project
git clone https://github.com/<username>/budget-tracker.git

cd budget-tracker

# Install dependencies
go mod tidy

# Run the server
go run main.go
