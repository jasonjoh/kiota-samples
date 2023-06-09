namespace RepoStars.Client.User.Repos {
    /// <summary>The default value for a squash merge commit title:- `PR_TITLE` - default to the pull request&apos;s title.- `COMMIT_OR_PR_TITLE` - default to the commit&apos;s title (if only one commit) or the pull request&apos;s title (when more than one commit).</summary>
    public enum ReposPostRequestBody_squash_merge_commit_title {
        PR_TITLE,
        COMMIT_OR_PR_TITLE,
    }
}
