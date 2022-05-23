# default postgres://postgres:password@localhost:5432/example?sslmode=disable
if [ "$DB_CONNECTION" == "" ]; then
  DB_CONNECTION="postgres://postgres:12345678@localhost:5432/ukraine-picture?sslmode=disable"
fi

function create_migration() {
  read -p "Enter migration name: " migration_name

  migrate create -ext sql -dir ../migrations -seq "$migration_name"
}

function run_migrations_up() {
  echo "Running migrations up..."

  migrate -database ${DB_CONNECTION} -path ../migrations up
}

function run_migrations_down() {
  echo "Running migrations down..."

  migrate -database ${DB_CONNECTION} -path ../migrations down
}

function force_migrations() {
  read -p "Enter version: " version

  migrate -path ../migrations -database ${DB_CONNECTION} force $version
}

echo "Migrations"
echo "-------------------------------------------"

if [ "$option" == "" ]; then
  echo "Choose your option:"
  echo "1 - run migrations up"
  echo "2 - run migrations down"
  echo "3 - create migration"
  echo "4 - force migration"
  echo "0. to exit!"
  read -p ": " option
fi

if [ "$option" == "1" ]; then
  run_migrations_up
elif [ "$option" == "2" ]; then
  run_migrations_down
elif [ "$option" == "3" ]; then
  create_migration
elif [ "$option" == "4" ]; then
  force_migrations
else
  echo "bye!"
  exit 0
fi
