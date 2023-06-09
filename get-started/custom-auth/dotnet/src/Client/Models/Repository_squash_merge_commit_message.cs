namespace RepoStars.Client.Models {
    /// <summary>The default value for a squash merge commit message:- `PR_BODY` - default to the pull request&apos;s body.- `COMMIT_MESSAGES` - default to the branch&apos;s commit messages.- `BLANK` - default to a blank commit message.</summary>
    public enum Repository_squash_merge_commit_message {
        PR_BODY,
        COMMIT_MESSAGES,
        BLANK,
    }
}
