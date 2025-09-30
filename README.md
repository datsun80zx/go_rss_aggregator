# Gator - RSS Feed Aggregator

Gator is a command-line RSS feed aggregator built in Go. It allows you to track multiple RSS feeds, follow specific feeds, and browse recent posts from your favorite sources. Think of it as your personal news aggregator that runs right from your terminal.

## Features

- **User Management**: Create and switch between multiple user accounts
- **Feed Management**: Add, follow, and unfollow RSS feeds
- **Automatic Aggregation**: Continuously fetch new posts from feeds at specified intervals
- **Post Browsing**: View recent posts from feeds you follow
- **PostgreSQL Storage**: All your feeds and posts are stored reliably in a database

## Prerequisites

Before you can run Gator, you'll need to have the following installed on your system:

### 1. Go (version 1.23 or later)
Go is the programming language Gator is built with. You'll need it to compile and install the program.

- **Download Go**: Visit [https://golang.org/dl/](https://golang.org/dl/) and follow the installation instructions for your operating system
- **Verify installation**: Run `go version` in your terminal. You should see something like `go version go1.23.4 linux/amd64`

### 2. PostgreSQL
PostgreSQL is the database that stores all your feeds, users, and posts.

- **Download PostgreSQL**: Visit [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
- **Installation guides**:
  - **macOS**: Consider using Homebrew: `brew install postgresql`
  - **Ubuntu/Debian**: `sudo apt-get install postgresql postgresql-contrib`
  - **Windows**: Use the interactive installer from the PostgreSQL website
- **Verify installation**: Run `psql --version` in your terminal

After installing PostgreSQL, you'll need to create a database for Gator:

```bash
# Connect to PostgreSQL as the default user
psql -U postgres

# Create a new database (replace 'gatordb' with your preferred name)
CREATE DATABASE gatordb;

# Exit psql
\q
```

## Installation

Once you have Go and PostgreSQL set up, installing Gator is straightforward. The `go install` command will download the code, compile it, and place the binary in your Go bin directory (usually `~/go/bin` or `$GOPATH/bin`).

```bash
go install github.com/datsun80zx/go_rss_aggregator.git@latest
```

After installation, make sure your Go bin directory is in your PATH. Add this line to your shell configuration file (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export PATH=$PATH:$HOME/go/bin
```

Then reload your shell configuration:

```bash
source ~/.bashrc  # or ~/.zshrc, depending on your shell
```

You should now be able to run the `gator` command from anywhere in your terminal!

## Configuration

Gator uses a JSON configuration file to store your database connection and current user information. This file needs to be created at `~/.gatorconfig.json` before you can use the program.

### Setting up the config file

Create the file `~/.gatorconfig.json` with your database connection details:

```json
{
  "db_url": "postgresql://username:password@localhost:5432/gatordb?sslmode=disable",
  "current_user_name": ""
}
```

Replace the connection string components:
- `username`: Your PostgreSQL username (often `postgres`)
- `password`: Your PostgreSQL password
- `localhost`: Database host (keep as is for local database)
- `5432`: PostgreSQL port (default is 5432)
- `gatordb`: The database name you created earlier

### Database Migrations

Before using Gator for the first time, you'll need to set up the database schema. If you have `goose` installed, you can run the migrations:

```bash
goose -dir sql/schema postgres "your-connection-string" up
```

## Using Gator

Gator is designed to be simple and intuitive. Here's how to get started and use the main features:

### Getting Started

First, create a new user account:

```bash
gator register yourname
```

This creates a new user and automatically logs you in. You can see all registered users with:

```bash
gator users
```

### Managing Feeds

#### Adding a new feed
To add an RSS feed to the database and automatically follow it:

```bash
gator addfeed "Tech News" "https://example.com/rss"
```

The first argument is a friendly name for the feed, and the second is the RSS feed URL.

#### Viewing all feeds
See all feeds that have been added to the database:

```bash
gator feeds
```

#### Following a feed
If a feed already exists in the database but you're not following it:

```bash
gator follow "https://example.com/rss"
```

#### Viewing your followed feeds
See which feeds you're currently following:

```bash
gator following
```

#### Unfollowing a feed
Stop following a feed (it remains in the database for other users):

```bash
gator unfollow "https://example.com/rss"
```

### Aggregating Posts

The aggregation command continuously fetches new posts from all feeds in the database. This runs in the foreground and updates at the interval you specify:

```bash
gator agg 30s  # Fetch feeds every 30 seconds
gator agg 5m   # Fetch feeds every 5 minutes
gator agg 1h   # Fetch feeds every hour
```

While the aggregator is running, it will:
- Check each feed for new posts
- Store new posts in the database
- Skip posts that have already been saved
- Mark when each feed was last fetched

You'll typically want to run this in a separate terminal window or as a background service.

### Browsing Posts

View recent posts from the feeds you follow:

```bash
gator browse      # Shows 2 most recent posts (default)
gator browse 5    # Shows 5 most recent posts
gator browse 10   # Shows 10 most recent posts
```

Posts are sorted by publication date, with the newest posts shown first.

### User Management

#### Switching users
If you have multiple user accounts, you can switch between them:

```bash
gator login differentuser
```

#### Resetting the database
To clear all users, feeds, and posts (use with caution!):

```bash
gator reset
```

## How It Works

Gator follows a simple but effective workflow:

1. **Configuration**: The `.gatorconfig.json` file stores your database connection and tracks which user is currently logged in

2. **User Context**: Most commands (like adding feeds or browsing) operate in the context of the currently logged-in user

3. **Feed Storage**: When you add a feed, it's stored in the database and available for all users to follow

4. **Aggregation**: The `agg` command runs continuously, checking each feed in sequence and storing new posts

5. **Browsing**: When you browse posts, Gator shows you only posts from feeds you follow, sorted by recency

## Architecture Notes

Gator demonstrates several Go best practices:

- **Modular design**: Commands, configuration, and database logic are separated into packages
- **SQL migrations**: Database schema is versioned using SQL migration files
- **Type safety**: Uses sqlc to generate type-safe database code from SQL queries
- **Command pattern**: Each CLI command is implemented as a separate handler function
- **Middleware pattern**: Authentication checks are implemented as middleware for commands that require a logged-in user

## Troubleshooting

### "Command not found" after installation
Make sure your Go bin directory is in your PATH. Run `echo $PATH` to check, and add the Go bin directory if it's missing.

### Database connection errors
- Verify PostgreSQL is running: `systemctl status postgresql` (Linux) or check your system's service manager
- Check your connection string in `~/.gatorconfig.json`
- Ensure the database exists and you have the right permissions

### "No user is currently logged in"
Some commands require you to be logged in. Run `gator login username` or `gator register newuser` first.

## Development

If you want to modify Gator or contribute to the project:

```bash
# Clone the repository
git clone https://github.com/datsun80zx/go_rss_aggregator.git
cd go_rss_aggregator

# Install dependencies
go mod download

# Run in development mode
go run . register testuser
go run . feeds

# Build the binary
go build -o gator .

# Install your local version
go install .
```

## License

[Add your license information here]

## Contributing

[Add contribution guidelines if you want others to contribute]