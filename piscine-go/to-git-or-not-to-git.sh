

URL="https://learn.zone01oujda.ma/assets/superhero/all.json"



SUBJECT_ID=170


curl -s "$URL" | jq -r --arg id "$SUBJECT_ID" '
    .[] | select(.id == ($id | tonumber)) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"
'