# How to contribute

We definitely welcome your patches and contributions to go-daraja! Please read the  [contribution guidelines](https://github.com/silaselisha/go-daraja/blob/master/CONTRIBUTING.md) before proceeding.

If you are new to github, please start by reading [Pull Request howto](https://help.github.com/articles/about-pull-requests/)


## Guidelines for Pull Requests
How to get your contributions merged smoothly and quickly.

- Create **small PRs** that are narrowly focused on **addressing a single
  concern**. We often times receive PRs that are trying to fix several things at
  a time, but only one fix is considered acceptable, nothing gets merged and
  both author's & review's time is wasted. Create more PRs to address different
  concerns and everyone will be happy.

- If you are searching for features to work on, issues are well-documented and usually   can beresolved with a single pull request.


- For speculative changes, consider opening an issue and discussing it first. 

- Provide a good **PR description** as a record of **what** change is being made
  and **why** it was made. Link to a github issue if it exists.

- If you want to fix formatting or style, consider whether your changes are an 
  obvious improvement or might be considered a personal preference. If a style 
  change is based on preference, it likely will not be accepted. If it corrects 
  widely agreed-upon anti-patterns, then please do create a PR and explain the 
  benefits of the change.

- Unless your PR is trivial, you should expect there will be reviewer comments
  that you'll need to address before merging. We'll mark it as `Status: Requires
  Reporter Clarification` if we expect you to respond to these comments in a
  timely manner. If the PR remains inactive for 6 days, it will be marked as
  `stale` and automatically close 7 days after that if we don't hear back from
  you.

- Maintain **clean commit history** and use **meaningful commit messages**. PRs
  with messy commit history are difficult to review and won't be merged. Fork the project
  and create a new branch to start woring on.

- Keep your PR up to date with upstream/master (if there are merge conflicts, we
  can't really merge your change).

- **All tests need to be passing** before your change can be merged. We
  recommend you **run tests locally** before creating your PR to catch breakages
  early on.
  - `go test --cover -v ./...` to run the tests

- Exceptions to the rules can be made if there's a compelling reason for doing so.
