# proxy-server
Description

Write HTTP server for proxying HTTP-requests to 3rd-party services.
The server is waiting HTTP-request from client (curl, for example). In request's body there should be message in JSON format. For example:

{
    "method": "GET",
    "url": "http://google.com",
    "headers": {
        "Authentication": "Basic bG9naW46cGFzc3dvcmQ=",
        ....
    }
}

Server forms valid HTTP-request to 3rd-party service with data from client's message and responses to client with JSON object:

{
    "id": <generated unique id>,
    "status": <HTTP status of 3rd-party service response>,
    "headers": {
        <headers array from 3rd-party service response>
    },
    "length": <content length of 3rd-party service response>
}

Server should have map to store requests from client and responses from 3rd-party service.

Usage

1) build docker image: 
	sudo docker build --tag proxy2 .
2) run docker container:
	sudo docker run -it -p 8080:8080 proxy


This application works as proxy-server redirected to another 3rd-party. 

The are 2 endpoints:

1) GET http://localhost:8080/proxy
Example request:
{
    "method": "GET",
    "url": "https://translate.yandex.com",
    "headers": {
        "content-type": "text/html"
    }
}
Example response:
{
    "id": "15161464-6ee8-4e88-5221-8f6b62991a8d",
    "status": "200 OK",
    "headers": {
        "Content-Length": [
            "8628"
        ],
        "Content-Type": [
            "text/html"
        ],
        "Strict-Transport-Security": [
            "max-age=31536000"
        ]
    },
    "length": 8628
}

2) GET http://localhost:8080/proxy
Example response:
[
    {
        "proxy_request": {
            "method": "GET",
            "url": "https://translate.yandex.com",
            "headers": {
                "content-type": "text/html"
            }
        },
        "proxy_response": {
            "id": "15161464-6ee8-4e88-5221-8f6b62991a8d",
            "status": "200 OK",
            "headers": {
                "Content-Length": [
                    "8628"
                ],
                "Content-Type": [
                    "text/html"
                ],
                "Strict-Transport-Security": [
                    "max-age=31536000"
                ]
            },
            "length": 8628
        }
    }
]


Also service has simple unit test for check up http statuses of services
