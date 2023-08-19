vim.api.nvim_open_win(0, true, {
    relative = "editor",
    title = "Example Window",
    title_pos = "center",
    width = 40,
    height = 20,
    row = 10,
    col = 25,
    style = "minimal",
    border = "single",
})
