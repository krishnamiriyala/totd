- Use `gd` to jump to definition of the word under cursor
- Use `Ctrl+o` to jump back to the previous position
- Use `dd` to delete current line
- Use `dt<char|space>` to delete from the cursor to next char or space provided
- Use `dw` delete the word under the cursor
- Press `i` to enter insert mode before the cursor
- Press `I` to enter insert mode at the beginning of the current line
- Press `o` to open a new line below the current line and enter insert mode
- Press `O` to open a new line above the current line and enter insert mode
- Understand the different modes: Normal, Insert, and Visual.
- Navigate with `h`, `j`, `k`, `l` for left, down, up, and right.
- Press `i` to enter Insert mode before the cursor.
- Use `w` to jump forward by word, and `b` to jump backward.
- Press `x` to delete the character under the cursor.
- Use `dw` to delete a word.
- Press `u` to undo, and `Ctrl + r` to redo changes.
- Yank text with `y`. For example, `yw` yanks a word.
- Use `p` to paste the yanked text.
- Search forward with `/`, backward with `?`, and `n` to find next.
- Use `:%s/old/new/g` to replace all occurrences of 'old' with 'new'.
- Save with ` :w`, quit with ` :q`. Force quit with ` :q!` if unsaved.
- Split horizontally with ` :split`, vertically with ` :vsplit`.
- Navigate splits with `Ctrl + w` followed by an arrow key.
- Use registers with `"ayw` to yank a word to register 'a', and `"ap` to paste from 'a'.
- Jump to the beginning of a line with `0`, and to the end with `$`.
- Use `gg` to go to the top of the file, and `G` to go to the bottom.
- Repeat the last change with `.`, and repeat the last search with `n`.
- Indent a block of lines with `>`, and unindent with `<` in Visual mode.
- Toggle line numbers with ` :set nu` and ` :set nonu`.
- Execute a shell command with ` :!command`. For example, ` :!ls`.
- Mark a location with `ma` and return to it with `'a`.
- Auto-indent the entire file with `gg=G`.
- Create a new line above the cursor with `O`, and below with `o` in Normal mode.
- To replace interactively, use `:%s/old/new/gc` and confirm each change.
- Add `i` after the first delimiter to make the search and replace case-insensitive: `:%s/old/new/gi`.
- Use `n` to jump to the next match in search, and `N` to go to the previous one.
- To search only in the selected range, use `:'<,'>s/old/new/g` after selecting in Visual mode.
- Use backreferences in the replacement: `:%s/\(pattern\)/\1_replacement/g`.
- To replace in a specific range of lines, use ` :10,20s/old/new/g` for lines 10 to 20.
- Use `&` in the replacement to repeat the search pattern: `:%s/pattern/&amp;/g`.
- Search with `* or #` to highlight the word under the cursor forward or backward.
- Use `/` followed by \(\) to search for multiple alternatives. For example, `/apple\|orange`.
- To match the whole word, use `\< and \>`. For example, `/\<word\>`.
- Highlight search results with ` :set hlsearch` and turn off with ` :nohlsearch`.
- Use ` :set incsearch` for incremental search highlighting as you type the pattern.
- Limit search to whole words with ` :set smartcase` and then use `/word` for case-sensitive or `/WORD` for case-insensitive search.
- Use ` :set ignorecase` to make searches case-insensitive.
- Search in all open buffers with ` :bufdo /pattern`.
- Use ` :vimgrep /pattern/ **/*.txt` to search in multiple files.
- Navigate through search results with ` :cn` for next and ` :cp` for previous.
