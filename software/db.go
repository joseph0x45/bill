package main

const dbSchema = `
  create table products (
    id, label, price
  )
  values (
    :id, :label, :price
  );
`

func runMigration() {

}
