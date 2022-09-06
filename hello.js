
conn = new Mongo();
db = conn.getDB("my_database");

//printjson(use Jaipur)
printjson(db.temp.findOne())
printjson(db.temp.insert({"city": "Jaipur",
"state": "Rajasthan"}))