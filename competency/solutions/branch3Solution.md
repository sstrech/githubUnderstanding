Solution to branch 3 (rebaseConflict)

This branch has 2 separate branches,  rebaseConflict and rebaseConflict2. Each branch edited the same part of function f3 which will create a rebase conflict when merging commences. After starting the rebase, since there is a conflict, the user will have to determine which part is prioritized. Something to point out here is that the branch that is in the rebase command will be treated as first and the branch that is checked out will be added on top of the other,
Example:
	git checkout rebaseConflict
	git rebase rebaseConflict2
	since there is a rebase conflict, the user needs to fix the conflict. Fix the 1 conflict from the files.
	git commit
	(for this specific example, it pushed the rebaseConflict2 branch before the rebaseConflict branch since it was rebased. The rebaseConflict branch was added on top of the rebaseConflict2 branch)
