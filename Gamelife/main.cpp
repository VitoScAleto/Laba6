#include "header.h"

using namespace std;


void task3() // X ИГРА В ЖИЗНЬ
{
    char action;
    int sizeMatrix;
    pair<int, int> lineRow;

    int line = 101;
    int column = 151;

    vector<vector<int>> playingField(line, vector<int>(column, 0));

    cout << "1-random 2-train q = exit" << endl;
    cout << "Enter action: ";
    cin >> action;

    switch (action)
    {
    case '1':
        generatePlayingField(playingField, line, column);

        playLife(playingField, line, column);
        break;

    case '2':
        lineRow = readingPlayingField(playingField);

        // cout << playingField << endl;
        // cout << lineRow.first << " " << lineRow.second;
        playLife(playingField, lineRow.first, lineRow.second);
        break;

    case 'q':
        return;
    default:
        cout << "unknown team\n";
    }
}