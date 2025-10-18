Solution to branch 1 (badCommitMessages)

To fix this branch, all that needs to happen is adding relevant information into the commit message. Each commit message should have a relevant title, and then a body that contains a few items. These items are “What is different now?,” “Why make this change?,” and “Any problems to look out for?”.
In this specific branch of badCommitMessages, if the following commit is to add a new function f4, then this will have to be communicated through the commit message.

Example: (after adding function f4)
	git checkout badCommitMessages
	git add competency/scenarios/main.go
	git commit 			…(in message ->)
	Added function f4 for main

	This commit added the function f4 that is similar to f3 which takes in a string and returns a string. This change was needed to provide more detail for the program. There is nothing notable to look out for problem-wise.
