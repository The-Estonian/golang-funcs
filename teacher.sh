location=$(find -name "*mystery*")
interview_number=$(grep "SEE INTERVIEW" $location/streets/Buckingham_Place |  cut -f 2 -d "#")
echo $interview_number
find . -name "interview-$interview_number" -exec cat {} +
echo $MAIN_SUSPECT
