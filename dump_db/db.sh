PASSWORD=dummyPass
LIMIT=60
NEXT_WAIT_TIME=0
until echo '0' | arangoinspect --server.ask-jwt-secret false --server.username "root" --server.password "${PASSWORD}" --quiet true | grep "dignostics collected" || [ $NEXT_WAIT_TIME -eq $LIMIT ]; do
   sleep 1
   echo "retry init arangodb: $(( ++NEXT_WAIT_TIME ))"
done

if [ $NEXT_WAIT_TIME -eq $LIMIT ]; then
  echo "

  ArangoDB not available!
  Initial script don't run

  "
  exit
fi

echo 'var list = db._databases(); if (list.includes("hash-db")) {db._dropDatabase("hash-db");} db._createDatabase("hash-db");' | arangosh --server.password ${PASSWORD}
arangosh --server.password ${PASSWORD} --server.database hash-db --javascript.execute-string 'db._drop("product-collection"); db._create("product-collection", {keyOptions: { type: "autoincrement"}});'
arangosh --server.password ${PASSWORD} --server.database hash-db --javascript.execute-string 'db._drop("user-collection"); db._create("user-collection", {keyOptions: { type: "autoincrement"}});'
arangoimp --file /opt/tools/products.json --collection product-collection --create-collection true --server.database hash-db --server.password ${PASSWORD}
arangoimp --file /opt/tools/users.json --collection user-collection --create-collection true --server.database hash-db --server.password ${PASSWORD}