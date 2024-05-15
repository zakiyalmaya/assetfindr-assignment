# AssetFindr-Assignment
AssetFindr Backend Engineer Assignment - API Development

This application is build using clean architecture and the following tech stack:
- Language: **Golang**
- Framework: **Gin**
- ORM: **GORM**
- Database: **PostgreSQL**

`config.go` is a configuration file that contains credential database values used by the application.

## Getting Started

- Clone the git repository
  ```sh
  https://github.com/zakiyalmaya/assetfindr-assignment.git
  ```

- Run the `.\main.go` file using this command
  ```sh
  go run .\main.go
  ```

## REST API

1. **Create Post**

    `POST /api/posts`

    ```sh
    curl --location 'http://localhost:8000/api/posts' \
    --header 'Content-Type: application/json' \
    --data '{
        "title": "Lorem Ipsum Dolor Sit Amet",
        "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
        "tags": ["Lorem", "Ipsum"]
    }'
    ```

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | title | string | Y | title of the post | 
        | content | string | Y | content of the post |
        | tags | array of string| Y | list of label tags |

        example:
        ```sh
        {
            "title": "Lorem Ipsum Dolor Sit Amet",
            "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
            "tags": ["Lorem", "Ipsum"]
        }
        ```

    - Response Body

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "Success"
        }
        ```

2. **Get All Post**

    `GET /api/posts`

    ```sh
    curl --location 'http://localhost:8000/api/posts'
    ```

    - Response Body

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "Success",
            "data": [
                {
                    "id": 1,
                    "title": "Lorem Ipsum Dolor Sit Amet",
                    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
                    "tags": [
                        {
                            "id": 1,
                            "label": "Lorem"
                        },
                        {
                            "id": 2,
                            "label": "Ipsum"
                        }
                    ]
                }
            ]
        } 
        ```

3. **Get Post By ID**
    
    `GET /api/posts/{id}`

    ```sh
    curl --location 'http://localhost:8000/api/posts/1'
    ```

    - Request Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | id | integer | Y | id of the post |

    - Request Body

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "Success",
            "data": {
                "id": 1,
                "title": "Lorem Ipsum Dolor Sit Amet",
                "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
                "tags": [
                    {
                        "id": 1,
                        "label": "Lorem"
                    },
                    {
                        "id": 2,
                        "label": "Ipsum"
                    }
                ]
            }
        }
        ```

4. **Update Post**

    `PUT /api/posts/{id}`

    ```sh
    curl --location --request PUT 'http://localhost:8000/api/posts/1' \
    --header 'Content-Type: application/json' \
    --data '{
        "title": "Updated Title",
        "content": "Updated Content",
        "tags": ["Updated Tag 1", "Updated Tag 2"]
    }'
    ```

    - Request Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | id | integer | Y | id of the post to be updated |
    
    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | title | string | Y | title of the post to be updated | 
        | content | string | Y | content of the post to be updated |
        | tags | array of string| Y | list of label tags to be updated |

        example:
        ```sh
        {
            "title": "Updated Title",
            "content": "Updated Content",
            "tags": ["Updated Tag 1", "Updated Tag 2"]
        }
        ```
        
    - Response Body

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "Success"
        }
        ```

5. **Delete Post**

    `DELETE /api/posts/{id}`

    ```sh
    curl --location --request DELETE 'http://localhost:8000/api/posts/1' \
    -- data ''
    ```

    - Request Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | id | integer | Y | id of the post to be deleted |
    
    - Response Body

        ```sh
        HTTP/1.1 200 OK
        {
            "responseMessage": "Success"
        }
        ```
