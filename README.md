# lets-work-together-my-theresa

## Description
We want you to implement a REST API endpoint that given a list of products, applies some
discounts to them and can be filtered.
You are free to choose whatever language and tools you are most comfortable with.
Please add instructions on how to run it and publish it in Github.

## Conditions 

  The prices are integers for example, 100.00€ would be 10000 .
  
  [ DONE ] You can store the products as you see fit (json file, in memory, rdbms of choice)
  
  Given that:
  [ DONE ] Products in the boots category have a 30% discount.
  [ DONE ] The product with sku = 000003 has a 15% discount.
  [ DONE ] Provide a single endpoint
  
  GET /products
  [ DONE ] Can be filtered by category as a query string parameter
  [ DONE ] (optional) Can be filtered by price as a query string parameter, this filter applies before discounts are applied.
  [ DONE ] Returns a list of Product with the given discounts applied when necessary
   Product model
  [ DONE ] price.currency is always EUR
  
  [ DONE ] When a product does not have a discount, price.final and price.original should be the same number and discount_percentage should be null.
  
  [ DONE ] When a product has a discount price.original is the original price, price.final is the amount with the discount applied and discount_percentage represents the
  applied discount with the % sign.
  
  Example product with a discount of 30% applied: 
  
    `{
      "sku": "000001",
      "name": "BV Lean leather ankle boots",
      "category": "boots",
      "price": {
          "original": 89000,
          "final": 62300,
          "discount_percentage": "30%",
          "currency": "EUR"
      }
    }`
  
  Example product without a discount:
  
      `{
        "sku": "000001",
        "name": "BV Lean leather ankle boots",
        "category": "boots",
        "price": {
            "original": 89000,
            "final": 89000,
            "discount_percentage": null,
            "currency": "EUR"
        }
      }`
      
## How to run this project 

I tried my best to make it fully automated, you only need Docker installed and be able to run make, then you can just `make run`

The system will inform you on the console when the service is ready, after that you can test it out through: 

`localhost:8080/products`
      
or apply one filter,

`localhost:8080/products?category=boots`

or maybe two? 

`localhost:8080/products?category=boots&price=89000`

## One doubt. 
In the test definition there is no rule on how to whenever there is more than one discount to apply, 
I took the liberty to choose the biggest discount possible. On real life I would ask that to be defined by the business instead.

## What have I done?

I tried my best to complete the task while delivering it using best practices, so you will find that I: 

* Used golang
* set up a layered architecture.
* made unit tests around domain, you can do it to by doing `go test` if you have go installed and want to try them out
* made asked api contract on the adapter layer.
* used protocol buffers! this service is also running on 8082 through GRPC, you can give it a try.
* Defined a simple domain that can be enhanced.
* used docker compose to set a runnable / reproduceable infrastructure.
* set up a configuration file 
* used a very simple "in memory" infrastructure.
* complied to the rules defined, at the domain, whenever the rules were about presentation I put them on the adapter
* Applied our takumi guidelines compendium available at www.github.com/takumi-software/guidelines. (like SOLID for example)

I hope you like it. I had a good time doing it, I know doing this for production is way more complex than this, and I am also 
know that many things can be improved, but time is gold. 

Total time Spent: 4:30Hrs.  

This repository will be opened to public until November 12th or until any notice from interested part to avoid it availability to others.