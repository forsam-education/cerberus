# Commit Convention

## Branch Typology

Branch name are composed as following :

```md
type/scope/purpose
```

For example, this add is of type doc, its scope is the commit, and its purpose is to explain how to name a branch.   
As a result, the branch name is `docs/contrib/branch-naming`

The purpose must be explicit and shouldn't lead to ask more questions that it is supposed to answer.

### Allowed type :
* **feat** : adding a new feature
* **fix** : fixing a bug
* **docs** : adding / fixing documentation

### Allowed scope:
* **api**
* **proxy**
* **contrib**

## Commit Typology

Make the changes to the code and tests and then commit to your branch. Be sure to follow the commit message conventions.

Commit message summaries must follow this basic format:

```none
Tag: Message (fixes #1234)
```

`Tag` should not be confused with git tag.
`Message` should not be confused with git commit message.

The `Tag` is one of the following:

- `Fix` - for a bug fix.
- `Update` - for a backwards-compatible enhancement.
- `Breaking` - for a backwards-incompatible enhancement.
- `Docs` - changes to documentation only.
- `Build` - changes to build process only.
- `New` - implemented a new feature.
- `Upgrade` - for a dependency upgrade.
- `Refacto` - for code refactoring.

The message summary should be a one-sentence description of the change. The issue number should be mentioned at the end. \* The commit message should say "(fixes #1234)" at the end of the description if it closes out an existing issue (replace 1234 with the issue number). If the commit doesn't completely fix the issue, then use `(refs #1234)` instead of `(fixes #1234)`.

Here are some good commit message summary examples:

```none
Build: Update Travis to only test Node 0.10 (refs #734)
Fix: Semi rule incorrectly flagging extra semicolon (fixes #840)
Upgrade: Esprima to 1.2, switch to using Esprima comment attachment (fixes #730)
```

The commit message format is important because these messages are used to create a changelog for each release. The tag and issue number help to create more consistent and useful changelogs.
