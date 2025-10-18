Solution to branch 2 (mergeConflict)

This branch has 2 separate branches, mergeConflict and mergeConflict2. Each branch edited the same part of function f3 which will create a merge conflict when merging commences. After starting the merge it will be up to the user to determine what should be added and what should not. For this specific case the user should choose the case from mergeConflict2 to override mergeConflict’s new f3.
Example:
	git checkout mergeConflict
	git merge mergeConflict2
	since there is a merge conflict, the user needs to fix the conflict. Fix the 2 conflicts from the files.
	git commit
