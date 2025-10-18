Solution to branch 5 (forcePushRecovery)

This branch is to show how to basically delete a commit or go back to a prior commit. This is important for when you accidentally included confidential information in a commit, but HAVE pushed it to the remote repository.
Example:
	git checkout forcePushRecovery
	git add competency/scenarios/main.go
	git commit
	git push origin forcePushRecovery
	(start recovery:)
	git reset --hard HEAD~1
	git push --force origin forcePushRecovery
