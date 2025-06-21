# 🛳️ Fileship
Fileship is a production-grade, full-stack file sharing platform — inspired by Dropbox, designed to teach real-world backend engineering. Built with cutting-edge technologies like **Next.js**, **TypeScript**, and **Docker**, Fileship is cloud-ready and developer-first.

# ⚙️ Tech Stack
* **Frontend**: Next.js (TypeScript)
* **Backend**: Next.js API Routes (Serverless)
* **Containerization**: Docker, Docker Compose

# 🛠️ Prerequisites

⚠️ **NOTE for First-Time Users** Before proceeding, please ensure you:
* Have **Docker** and **Docker Desktop** installed and set up
* Keep **Docker Desktop running at all times** while using the project

# 🧾 Getting Started
## 🌀 Step 1: Clone the Repository

```
git clone https://github.com/YadavPrasad/fileship.git
cd fileship
```

## 🧱 Step 2: Build and Run the Full Stack Application

This will build everything and spin up all services:

```
docker compose up --build
```

🔁 This will:
* Build the frontend and backend services
* Spin up all containers
* Make the application available at `http://localhost:3000`

🚫 Note: The **frontend will NOT hot-reload** inside Docker

## 🧪 Optional: Development Mode

If you want faster development with hot-reload:

```
npm install
npm run dev
```

This runs the application locally outside Docker for quicker iteration.

## 🧹 To Stop and Clean Up

When you're done, gracefully shut everything down:

```
docker compose down
```

# 💡 Tips

* Ensure Docker Desktop is running before executing any `docker compose` commands
* For live frontend development, consider running the application locally outside Docker
* Use `docker compose logs` to view container logs for debugging

# 🧑‍💻 Author

**Yadav Prasad**
* GitHub: [@YadavPrasad](https://github.com/YadavPrasad)