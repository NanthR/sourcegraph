import { searchQueryValidator } from '../../../../../pages/insights/creation/capture-group/utils/search-query-validator'
import { ValidationResult } from '../../../../form'
import { createRequiredValidator, composeValidators } from '../../../../form/hooks/validators'

export const SERIES_NAME_VALIDATORS = createRequiredValidator('Name is a required field for data series.')

export const SERIES_QUERY_VALIDATORS = composeValidators([
    createRequiredValidator('Query is a required field for data series.'),
    (value: string | undefined): ValidationResult => {
        // TODO: decouple searchQueryValidator (do not use anything from capture group creation UI)
        const { isNotContext, isNotRepo } = searchQueryValidator(value)

        if (!isNotContext) {
            return 'The `context:` filter is not supported; instead, run over all repositories and use the `context:` on the filter panel after creation'
        }

        if (!isNotRepo) {
            return 'Do not include a `repo:` filter; add targeted repositories above, or filter repos on the filter panel after creation'
        }
    },
])
