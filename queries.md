# GraphQL test queries

Run these in the GraphQL Playground at `http://localhost:8080/playground` after
starting the services. Run the create mutations first, then copy the returned
IDs into the order mutation.

## Create an account

```graphql
mutation {
  createAccount(account: { name: "Test Account" }) {
    id
    name
  }
}
```

## List accounts

```graphql
query {
  accounts(pagination: { skip: 0, take: 10 }) {
    id
    name
  }
}
```

## Get one account by ID

Replace `ACCOUNT_ID` with an ID returned by `createAccount`.

```graphql
query {
  accounts(id: "ACCOUNT_ID") {
    id
    name
  }
}
```

## Create a product

```graphql
mutation {
  createProduct(
    product: {
      name: "Test Coffee"
      description: "Whole bean coffee"
      price: 12.5
    }
  ) {
    id
    name
    description
    price
  }
}
```

## List products

```graphql
query {
  products(pagination: { skip: 0, take: 10 }) {
    id
    name
    description
    price
  }
}
```

## Search products

```graphql
query {
  products(query: "coffee", pagination: { skip: 0, take: 10 }) {
    id
    name
    description
    price
  }
}
```

## Get one product by ID

Replace `PRODUCT_ID` with an ID returned by `createProduct`.

```graphql
query {
  products(id: "PRODUCT_ID") {
    id
    name
    description
    price
  }
}
```

## Create an order

Replace `ACCOUNT_ID` and `PRODUCT_ID` with IDs returned by the create mutations.

```graphql
mutation {
  createOrder(
    order: {
      accountId: "ACCOUNT_ID"
      products: [{ id: "PRODUCT_ID", quantity: 2 }]
    }
  ) {
    id
    createdAt
    totalPrice
  }
}
```

## Get an account and its orders

Replace `ACCOUNT_ID` with the account ID used to create the order.

```graphql
query {
  accounts(id: "ACCOUNT_ID") {
    id
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        id
        name
        description
        price
        quantity
      }
    }
  }
}
```
