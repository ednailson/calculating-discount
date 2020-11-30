import arangojs from "arangojs"

const db = arangojs({
    databaseName: "hash-db",
    url: "http://arangodb.service.com.br:8529",
    auth: {
        username: "root",
        password: "dummyPass"
    }
})

export default db