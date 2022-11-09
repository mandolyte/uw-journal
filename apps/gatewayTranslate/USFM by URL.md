# USFM by URL

Goal: a translator will be able to view and/or edit USFM text obtained via URL.

There are three general use cases.

## Use Case 1

The USFM text comes from DCS, but user does not have access to save the changes. For example, the Original Languages Greek and Hebrew texts, wherein the user, unless thay are a member of the unfoldingWord organization, will not be able to save any changes. This use case is expected to common for any translator who know Greek or Hebrew.

## Use Case 2

The USFM text comes from DCS and the user has access to save the changes, since they are a member of the owning organization. This would be a common use case for a translator who wanted to refer to legacy texts stored on DCS. This would also be the use case for making changes to the Greek or Hebrew texts.

## Use Case 3

The USFM text does not come from DCS. This use case might be of use to anyone who wished to use our tools. Since the edits could not be saved, gT will need to provide a download option for the edited text.

## Note on Original Languages

Since Greek and Hebrew is anticipated to be common use case, gT should conside a built in way to obtain the original language texts. 

## Considerations and Constraints

1. If a DCS text is from a "release" URL, then even if the user has edit rights, we should only permit edits to be done from URLs from the master branch.
2. The book ID must be a valid one. ([reference](https://github.com/Copenhagen-Alliance/versification-specification/blob/master/versification-mappings/standard-mappings/eng.json))
3. For DCS texts, inspection of the URL will yield both an organization and a repo. These might be used to formulate the required doc set ID.
4. For non-DCS texts, creation of the required doc set ID will take some thought.
5. Since we cannot always determine the book ID from the URL, consider forcing the user to specify the book from the drop down list. If we wish to support apocryphal or other non-canonical texts, our drop down should contain all the books from the reference in item 2 above. 
6. By requiring the user to supply the book ID, then these texts will be able to participate in the linked card navigation and synchronization. (see [issue 16](https://github.com/unfoldingWord/gateway-translate/issues/16))

