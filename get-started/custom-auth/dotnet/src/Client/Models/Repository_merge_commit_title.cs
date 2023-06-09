namespace RepoStars.Client.Models {
    /// <summary>The default value for a merge commit title.- `PR_TITLE` - default to the pull request&apos;s title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).</summary>
    public enum Repository_merge_commit_title {
        PR_TITLE,
        MERGE_MESSAGE,
    }
}
