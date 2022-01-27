package gh

// Graphql API struct
type Graphql struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

// Error of the Graphql API
type Error struct {
	Message string `json:"message"`
}

// StatusData struct from the Graphql API response json
type StatusData struct {
	Data struct {
		User struct {
			Status Status `json:"status"`
		} `json:"user"`
	} `json:"data"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

// UserData struct from the Graphql API response json
type UserData struct {
	Data struct {
		Viewer User `json:"viewer"`
	} `json:"data"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

// User info struct
type User struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}

// Status struct
type Status struct {
	Emoji   string `json:"emoji"`
	Message string `json:"message"`
}

// GetUserData is the Graphql query to get the user data
const GetUserData = `{
  viewer {
    login
    name
    id
  }
}`

// SetUserStatusQuery is the GraphQL query to set the user status
const SetUserStatusQuery = `mutation($emoji: String!, $message: String!) {
  changeUserStatus(input: {
	  emoji: $emoji
	  message: $message
  }) {
	status {
	  emoji
	  message
	}
  }
}`

// ClearUserStatusQuery is the GraphQL query to clear the user status
const ClearUserStatusQuery = `mutation {
  changeUserStatus(input: {}) {
    status {
	  message
    }
  }
}`

// GetUserStatusQuery is the GraphQL query to get the user status
const GetUserStatusQuery = `{
  user(login: %q) {
    status {
      emoji
      message
    }
  }
}`
