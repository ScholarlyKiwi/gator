# Gator RSS Aggregator

Gator is a Multi-user RSS Feed Aggregator project created as a guided project for boots.dev.

Usage:

gator (command) (aruguments)

Commands:

addfeed (feed title) (feed url)     - adds and follows a feed for the current user
agg (duration)                      - retrieves all posts for the current user every duration
browse (optional limit)             - displays the latest posts for the user to the limit (default 2)
feeds                               - lists all feeds stores in gator
follow (feed url)                   - follows a feed for the current user, adding it if it has not been retrieved
following                           - lists all feeds followed by the current user
login (user)                        - changes the current user
register (user)                     - creates a new user and logs in as the user
reset                               - delete all users
unfollow (feed url)                 - stops follows a feed for the current user
users                               - list all users
