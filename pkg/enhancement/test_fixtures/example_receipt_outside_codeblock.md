## Enhancements Tracking Introduction

This is a space for someone shepherding an enhancement through the release process
to coordinate with the release team to verify that

- [x] the enhancement is part of a SIG sponsored KEP
- [x] the shepherd(s) are updating tests
- [x] the shepherd(s) are preparing release notes
- [x] the shepherd(s) are working on updating documentation
- [x] contact info is available during the milestone burndown

## Enhancements Tracking Launch Checklist

---
title: "A Great Idea, Delivered Incrementally"
authors:
- "justaugustus"
- "calebamiles"
sponsoring_sig: "SIG Release"
affected_subprojects:
- "release-engineering"
kep_location: "https://github.com/kubernetes/enhancements/blob/master/keps/0001a-meta-kep-implementation.md"
test_locations:
- "https://github.com/planctae/enhancements-tracking-ng/blob/master/pkg/enhancement/check_receipt_test.go"
- "https://github.com/planctae/enhancements-tracking-ng/blob/master/pkg/enhancement/extract_receipt_test.go"
- "https://github.com/planctae/enhancements-tracking-ng/blob/master/pkg/enhancement/check_receipt_test.go"
release_notes_location: "https://github.com/planctae/enhancements-tracking-ng/blob/master/README.md" # should be retrievable with http.Get()
documentation_locations:
- "https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/enhancement-authors.md"
- "https://github.com/planctae/enhancements-tracking-ng/blob/master/guides/release-team.md"
contact_email: "kubernetes-sig-release@googlegroups.com" # may be a "plus aliased" SIG mailing list
current_maturity: "RC" # must be one of RC|alpha|beta|GA
target_maturity: "alpha" # must be one mof RC|alpha|beta|GA

