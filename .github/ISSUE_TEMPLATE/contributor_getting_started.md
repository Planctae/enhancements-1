---
name: Getting Started with the Enhancement Tracking Process
about: Enhancement Contributor Workflow 
title: 'Help Needed with the Enhancement Tracking Process'
labels: 'support/contributor'
assignees: 'calebamiles'
---

## Enhancements Tracking Introduction

This is a space for someone shepherding an enhancement through the release process
to coordinate with the release team to verify that:

- [ ] the enhancement is part of a SIG sponsored KEP
- [ ] the shepherd(s) are updating tests
- [ ] the shepherd(s) are preparing release notes
- [ ] the shepherd(s) are working on updating documentation
- [ ] contact info is available during the milestone burndown

## Enhancements Tracking Receipt

Each accepted enhancement will correspond to an "Enhancements Tracking Receipt"
that is checked into this repository. In the future this receipt file will be
replaced with a pointer to KEP content and there will be no need for almost anyone
to produce a receipt manually.


With the introduction of [draft pull requests][], along with our existing system for
[marking a pull request as WIP][], it not unreasonable to expect that the requester
of enhancement tracking to coordinate with the sponsoring SIG, as well as SIG Docs,
to open draft pull requests before tracking by the Release Team begins. Currently, an
enhancement tracking receipt has the following form where **all fields are mandatory**:

```
---
title: ""
authors:
- ""
sponsoring_sig: ""
affected_subprojects:
- ""
kep_location: ""
test_locations:
- ""
release_notes_location: "" # should be retrievable with http.Get()
documentation_locations:
- ""
contact_email: "" # may be a "plus aliased" SIG mailing list
current_maturity: "" # must be one of RC|alpha|beta|GA
target_maturity: "" # must be one mof RC|alpha|beta|GA
```

[draft pull requests]: https://github.blog/2019-02-14-introducing-draft-pull-requests/
[marking a pull request as WIP]: https://prow.k8s.io/plugins

