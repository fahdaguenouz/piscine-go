InterviewNUM=$(sed -n '179p' ./streets/Buckingham_Place | sed 's/.*#//')
echo $InterviewNUM
cat $(echo "./interviews/interview-${InterviewNUM}")
echo $MAIN_SUSPECT