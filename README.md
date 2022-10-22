<a name="readme-top"></a>

<!-- GETTING STARTED -->
## Getting Started

This project is concentrated on reading big .csv files, storing them in MySQL database and getting them using their ID.

### Installation

_Installing and setting up app.

1. Clone the repo
   ```sh
   git clone https://github.com/BoroBalasan/Promotions.git
   ```
2. Install mysql driver packages
   ```sh
   go get github.com/go-sql-driver/mysql
   ```
3. Install Docker if you don't have on your local   
4. In terminal run this commands
   ```sh
   docker compose build --no-cache
   ```
   ```sh
   docker compose up
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

The app has very ugly front-end part with only one index.html which you can use to upload your .csv files

_Please use (localhost:1321) to upload your files_ . You can use Promotions\Resources\promotions.csv for test purposes .
_Please use (localhost:1321/promotions?={ID}) to see promotion information were ID is Id column of the promotion_

<p align="right">(<a href="#readme-top">back to top</a>)</p>
