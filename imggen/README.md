# Project Name
image generator


## Getting Started
To install and launch the project, you need to have Docker and Docker Compose installed on your system.

### Prerequisites

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Docker Compose: [Install Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository: https://github.com/anvme/gen-uploader.git
2. cd gen-uploader
3. Firstly: docker-compose up --build. Then with -d flag


### Api
    POST /api/post - create image and json with metainfo

### JSON Format

```json
{
  "tokenName": "SOLANA",
  "tokenSymbol": "QWE",
  "tokenAddr": "DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263",
  "tokenImg": "https://www.freepnglogos.com/uploads/lion-logo-png/commercial-real-estate-black-lion-investment-group-los-0.png",
  "NFTBackground": "bg1",
  "tariffs": [
    {
      "lock_period_days": "30",
      "apy": 5
    },
    {
      "lock_period_days": "90",
      "apy": 7
    },
    {
      "lock_period_days": "180",
      "apy": 10.5
    }
  ]
}
```


 
