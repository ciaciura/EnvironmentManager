db = db.getSiblingDB('your_database_name');

db.createUser({
  user: "default_user",
  pwd: "default_password",
  roles: [
    { role: "readWrite", db: "your_database_name" }
  ]
}); 