requestedName=$(curl https://01.kood.tech/assets/superhero/all.json | jq " .[] | select( .id == 170 ).name")
requestedPower=$(curl https://01.kood.tech/assets/superhero/all.json | jq " .[] | select( .id == 170 ).powerstats.power")
requestedGender=$(curl https://01.kood.tech/assets/superhero/all.json | jq " .[] | select( .id == 170 ).appearance.gender")
echo $requestedName 
echo $requestedPower 
echo $requestedGender
