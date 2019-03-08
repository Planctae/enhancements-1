## Terms

- `release-N`: the current release

### Activities Before Enhancements Freeze

Review the Pull Requests which are created against the `release-N/proposed`; these Pull Requests
will hopefully land in the "[on the platform][]" column of the [GitHub Project Board][]
automatically.

After reviewing a Pull Request for correctness, merge it into the `release-N/proposed`
subdirectory. Continuous Integration (CI) should hopefully signal when an enhancement is ready
for merge. It is important to merge _all_ pull requests into the `proposed` directory in order
to retain an accurate count of all the enhancements that have been considered by the enhancements
team.

Enhancement receipts which pass automated testing, but upon human review do not meet a
reasonable standard for completeness should be merged into the `release-N/rejected` subdirectory.
The commit message should clearly communicate the rationale for rejecting the enhancement.

Merge acceptable enhancements into the `release-N/accepted` directory. At this point file, or
request the proposer file, a new Pull Request which copies, not moves, the accepted enhancement
into the `release-N/shipped` directory: this pull request will be tracked throughout the release
cycle. Apologies for the busy work, but interesed parties can help change that through the
Power of Automation™.

[GitHub Project Board]: https://github.com/Planctae/enhancements-tracking-ng/projects/1
[on the platform]: https://github.com/Planctae/enhancements-tracking-ng/projects/1#column-4565553

### Implementing Enhancements Freeze

Issue a [blockade][] against the current `release-N/{proposed, accepted}` subdirectories.

[blockade]: https://prow.k8s.io/plugins

### Activities During Code Freeze

At this point in the release process there should only be Pull Requests open against the
`release-N/shipped` subdirectory. Pull Requests for enhancements that have not already been
marked as `accepted` should be rejected by merging them into the `release-N/rejected` subdirectory.
For any Pull Requests where the state is unclear move them into  the "[styx column][]" of
the project board. Why "styx"? Seems like a [fitting name][]. These Pull Requests should be
reviewed in a release burndown meeting for final resolution.

[styx column]: https://github.com/Planctae/enhancements-tracking-ng/projects/1#column-4565573
[fitting name]: https://en.wikipedia.org/wiki/Styx

### Running the Next Release Train

Rather than "lifting" the enhancements freeze, what we do is start the next release
train running. In the future we hope to run smaller trains more often but we need to
start conditioning people to expect a train to run on time first. Starting the next
release train involves:

1. Staffing the next Release Team Lead
1. Staffing the next Enhancements Lead
1. Adding the leads and team to `OWNERS_ALIAS`
1. Running `reltracker init release-<N+1>`; and merging the changes to this repository
1. Announcing that the train is boarding:

   - [kubernetes-dev][]
   - [Community Meeting][]
   - [SIG PM][]
   - [SIG Architecture][]

[kubernetes-dev]: https://groups.google.com/forum/#!forum/kubernetes-dev
[SIG PM]: https://github.com/kubernetes/community/tree/master/sig-pm
[SIG Architecture]: https://github.com/kubernetes/community/tree/master/sig-architecture
[Community Meeting]: https://github.com/kubernetes/community/blob/master/events/community-meeting.md

