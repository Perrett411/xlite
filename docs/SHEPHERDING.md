# Shepherding Protocol: Equality for All

Shepherding is how maintainers guide a contribution from first submission to resolution. This protocol ensures every contributor is treated equally, and that shepherding is additive: it never requires rewriting a contributor's work.

## Principles

1. **Equal treatment.** Every issue and pull request gets the same standards, response expectations, and tone, regardless of who the author is, their affiliation, tenure, background, or identity.
2. **No rewrites.** Maintainers do not rewrite, replace, or squash away a contributor's work without their consent. Feedback is given as review comments and suggestions; the author decides how to apply them.
3. **Attribution is preserved.** Original authorship is kept in commit history. Co-authors are credited with `Co-authored-by` trailers.
4. **Transparency.** Decisions and reasons are stated in the open, in the issue or pull request.

## Protocol

### 1. Acknowledge
- Every new issue or pull request receives a first response using the same triage process (see [triage.md](triage.md)).
- Responses are not prioritized by author identity.

### 2. Assign a shepherd
- A maintainer is assigned to guide the contribution, rotating fairly among available maintainers.
- The shepherd is named in the thread so the contributor knows who to ask.

### 3. Review without rewriting
- Use review comments and GitHub suggestions rather than pushing commits over the author's branch.
- Push to a contributor's branch only if they explicitly invite it.
- If the author becomes unavailable, ask before continuing their work. If a new pull request is needed, keep their commits and credit them.

### 4. Apply consistent criteria
Evaluate all contributions against the same published criteria:
- Fit with project scope and direction
- Tests and documentation
- Code style and the repository's contribution guidelines
- Security and compatibility

### 5. Resolve openly
- Accepted: merge with the author's commits and attribution intact.
- Declined: explain the reason, point to the criteria, and suggest alternatives where possible.
- Stalled: follow up politely, then close with a note that the work can be reopened.

## Conduct

All participants follow the project's [Code of Conduct](../.github/CODE-OF-CONDUCT.md). Discrimination or unequal treatment in shepherding should be reported to the maintainers and will be reviewed.
