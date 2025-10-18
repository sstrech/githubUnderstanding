Solution to branch 4 (recoverLostCommit)

This branch is to show how to basically delete a commit or go back to a prior commit. This is important for when you accidentally included confidential information in a commit, but have NOT pushed it to the remote repository yet.
Example:
	git checkout recoverLostCommit
	git add competency/scenarios/main.go
	git commit 	…(add a good message in editor)
(start recovery:)
	git reset --hard HEAD~1
