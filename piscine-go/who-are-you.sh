URL="https://learn.zone01oujda.ma/assets/superhero/all.json"

SUBJECT_ID=70

NAME=$(curl -s "$URL" | jq -r --arg id "$SUBJECT_ID" '.[] | select(.id == ($id | tonumber)) | .name')

echo "\"$NAME\""