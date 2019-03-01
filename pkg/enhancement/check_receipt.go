package enhancement

import (
	"errors"
	"fmt"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/maturity"
)

func CheckReceipt(receipt TrackingReceipt) error {
	var err error
	var errs *multierror.Error

	err = checkThatTitleIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatAuthorsAreGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatSponsoringSIGIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatAffectedSubprojectsAreGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatKEPLocationIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatTestLocationsAreGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatReleaseNoteLocationIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatDocumentationLocationsAreGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatContactEmailIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatCurrentMaturityIsGiven(receipt)
	errs = multierror.Append(errs, err)

	err = checkThatTargetMaturityIsGiven(receipt)
	errs = multierror.Append(errs, err)

	return errs.ErrorOrNil()
}

func checkThatTitleIsGiven(receipt TrackingReceipt) error {
	if title := receipt.Title(); title == "" {
		return errors.New("no title given")
	}

	return nil
}

func checkThatAuthorsAreGiven(receipt TrackingReceipt) error {
	if authors := receipt.Authors(); len(authors) == 0 {
		return errors.New("no authors given")
	}

	return nil
}

func checkThatSponsoringSIGIsGiven(receipt TrackingReceipt) error {
	if sponsoringSIG := receipt.SponsoringSIG(); sponsoringSIG == "" {
		return errors.New("no sponsoring SIG given")
	}

	return nil
}

func checkThatAffectedSubprojectsAreGiven(receipt TrackingReceipt) error {
	if affectedSubprojects := receipt.AffectedSubprojects(); len(affectedSubprojects) == 0 {
		return errors.New("no affected subprojects given")
	}

	return nil
}

func checkThatKEPLocationIsGiven(receipt TrackingReceipt) error {
	if kepLocation := receipt.KEPLocation(); kepLocation == "" {
		return errors.New("no KEP location given")
	}

	return nil
}

func checkThatTestLocationsAreGiven(receipt TrackingReceipt) error {
	if testLocations := receipt.TestLocations(); len(testLocations) == 0 {
		return errors.New("no test locations given")
	}

	return nil
}

func checkThatReleaseNoteLocationIsGiven(receipt TrackingReceipt) error {
	if releaseNoteLocation := receipt.ReleaseNoteLocation(); releaseNoteLocation == "" {
		return errors.New("no release note location given")
	}

	return nil
}

func checkThatDocumentationLocationsAreGiven(receipt TrackingReceipt) error {
	if documentationLocations := receipt.DocumentationLocations(); len(documentationLocations) == 0 {
		return errors.New("no documentation locations given")
	}

	return nil
}

func checkThatContactEmailIsGiven(receipt TrackingReceipt) error {
	if contactEmail := receipt.ContactEmail(); contactEmail == "" {
		return errors.New("no contact email given")
	}

	return nil
}

func checkThatMaturityIsValid(maturityLevel string) error {
	match := maturity.LevelMatches

	switch {
	case match(maturityLevel, maturity.ReleaseCandidate):
		return nil
	case match(maturityLevel, maturity.Alpha):
		return nil
	case match(maturityLevel, maturity.Beta):
		return nil
	case match(maturityLevel, maturity.GenerallyAvailable):
		return nil
	default:
		return fmt.Errorf("given state: %s is not valid", maturityLevel)
	}
}

func checkThatCurrentMaturityIsGiven(receipt TrackingReceipt) error {
	currentMaturity := receipt.CurrentMaturity()
	if currentMaturity == "" {
		return errors.New("no current maturity level given")
	}

	return checkThatMaturityIsValid(currentMaturity)
}

func checkThatTargetMaturityIsGiven(receipt TrackingReceipt) error {
	targetMaturity := receipt.TargetMaturity()
	if targetMaturity == "" {
		return errors.New("no target maturity level given")
	}

	return checkThatMaturityIsValid(targetMaturity)
}
