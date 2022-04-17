# Final Project of Picus Security Golang Bootcamp

- This project is a sample E-Commerce Back-End project that is coded in Go Programming Language. 

## Categories:
- There are 3 categories (Shoes,Pants,Glasses). They are created using csv files. You may find the files in csvfiles.
## Products:
- Products can be added by Super User manually using APIs or Super User can modify the csv files to change the products.

## User Types:
- There are 3 user types.
    - Super User
    - Customer
    - Guest

- Admin can be detected by username. I did not add any role because I wanted to keep it simple.
- An Admin can be also a customer. So, I think any user which have valid token, can buy something and add items to cart.
- Guests have no token because they do not to login the system. So they can just view the products.

## Used technologies:
- JSON Web Token
- Gin (To activate APIs)
- Postman (For check the APIs)
- PostgreSQL (Database which tables are handeled)
- GORM (To send values to Database)


## Usage
- Firstly, you need to create a super user to insert the products into the Database. Then you can create normal users. Admin authorezation can be activated if username is admin. Then you can use the commands on main.go file using Postman or terminal of your computer

## Challenging Topics:
- Expiration date of an order in 14 days  and pagination of product/user listing was a bit challenging. I am not sure if they work successfully. If you have any advice please share your ideas with me. 
