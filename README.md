# Package Counter App

Package Counter App is a simple Go application that calculates the minimum number of packs required to fulfill a customer's order based on predefined pack sizes. The application also provides an API for dynamic pack size management and a web-based user interface.

## Features

- Calculate the minimum number of packs needed to fulfill an order.
- Dynamically manage and adjust pack sizes.
- Web-based user interface for easy interaction.

## Getting Started

### Prerequisites

- Go installed on your system.
- Dependencies managed with Go Modules.

### Installation

1. Clone the repository:

   ```bash
   git clone 

    Navigate to the project directory:

    bash

cd package-counter-app

Run the application:

bash

    go run main.go

Usage

    Access the web-based user interface by opening a web browser and visiting http://localhost:5050.
    Enter the order quantity and click the "Calculate" button to see the minimum number of packs required.

API Endpoints



Calculate Packs:

bash

GET /api/v1/packs/calculate?quantity=<order_quantity>




Configuration

    The application uses the PACK_SIZES environment variable to load the initial pack sizes. You can modify this variable in the .env file.