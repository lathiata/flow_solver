var boardState = {
    'height' : 5,
    'width' : 5,
    'board' : [],
};

$(document).ready(function () {
    var $fadeDivsIn = $('body .container div.initial-animate');
    var fadeInTime = 0;
    var fadeInTimeDifference = 500;

    // Stagger initial fade in animation
    $fadeDivsIn.each((index, value) => {
        var $targetDiv = $(value);
        setTimeout(() => {
            $targetDiv.fadeIn();
        }, fadeInTime);
        fadeInTime += fadeInTimeDifference;
    });

    setTimeout(() => {
        initBoardState();
        renderBoard();
    }, fadeInTime);
});

var renderBoard = function () {
    var board = boardState.board;
    var html = '';
    for (var i = 0; i < board.length; i++)
    {
        html += "<div class='board-row'>"
        for (var j = 0; j < board[0].length; j++)
        {
            html += "<div class='board-cell " + board[i][j].color + "'></div>"
        }
        html += "</div>"
    }

    var $createBoard = $('body .container div#create-board');
    $createBoard.fadeOut(() => {
        $createBoard.html(html);
        $createBoard.fadeIn();
    });
};

var updateBoard = function () {
    var board = boardState.board;
    var rowDiff = boardState.height - boardState.board.length;
    if (rowDiff > 0)
    {
        // Add a row
        var row = [];
        for (var j = 0; j < boardState.width; j++)
        {
            row.push({
                'color' : ''
            });
        }
        board.push(row);
    }
    else if (rowDiff < 0)
    {
        // Delete a row
        board = board.slice(board.length - 2);
    }

    var columnDiff = boardState.width - boardState.board[0].length;
    if (boardState.width != boardState.board[0].length)
    {
        // Add a column
        for (var i = 0; i < board.length; i++)
        {
            board[i].push({
                'color' : ''
            });
        }
    }
    else if (columnDiff < 0)
    {
        // Delete a column
        for (var i = 0; i < board.length; i++)
        {
            board[i] = board[i].slice(board[0].length - 2);
        }
    }
};

var initBoardState = function () {
    for (var i = 0; i < boardState.height; i++)
    {
        var row = [];
        for (var j = 0; j < boardState.width; j++)
        {
            row.push({
                'color' : ''
            });
        }
        boardState.board.push(row);
    }
};
