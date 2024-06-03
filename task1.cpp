#include <iostream>
#include <vector>
#include <string>
#include<random>
#include <algorithm>
#include <iomanip>
#include <Windows.h>

using namespace std;

int GetRandomNumber(int min, int max)
{
    random_device rd;//random_device, который является источником недетерминированных случайных чисел.
    //Затем мы используем это устройство для заполнения генератора std::minstd_rand. Затем функция генератора() используется для генерации случайного числа
    minstd_rand generator(rd());

    uniform_int_distribution<int> distribution(min, max);// функция destribition для задания диапозона значений
    int random_number = distribution(generator);
    return random_number;
}

void  task1()
{
    cout << "Task 1" << endl;
    int quantityLine = 5, quantityColumns = 6;
    int maxSrednee = 0, numberLine = 0;
    
    vector<vector<int>> randomMatrix(quantityLine, vector<int>(quantityColumns));;
  
    for (auto& i : randomMatrix)
    {
        for (auto& j : i)
        {
            j = GetRandomNumber(0, 100);
        }
    }
    for (auto& i : randomMatrix)
    {
        for (auto& j : i)
        {
            cout << setw(4) << j;
        }
        cout << endl;
    }

    for (int i = 0; i < quantityLine; i++)
    {
        float sredneeArithmetic = 0;

        for (int j = 0; j < quantityColumns; j++)
        {
            sredneeArithmetic += randomMatrix[i][j];// среднее арифетическое строки 
        }
        if (maxSrednee < sredneeArithmetic)// максимальное среднее арифметическое
        {
            maxSrednee = sredneeArithmetic;
            numberLine = i;
        }
        cout << "Номер строки " << i << " среднее арифметическое = " << sredneeArithmetic / quantityColumns << endl;
    }
    cout << "Наибольшее среднеее арифметическое в строке " << numberLine;
    cout << endl;
}


void task2()
{
    cout << "Task 2" << endl;

    int sizeSquareMatrix = 15;
    
    vector<vector<int>> squareMatrix(sizeSquareMatrix, vector<int>(sizeSquareMatrix));

    for (auto& i : squareMatrix)
    {
        for (auto& j : i)
        {
            j = GetRandomNumber(-100, 100);
            cout << setw(4) <<j;
        }
        cout << endl;
    }

    for (int i = 0; i < sizeSquareMatrix; i++)// цикл наоборот. Берем первый элемент строки
    {
        vector<int> columns(sizeSquareMatrix);// вектор для хранение стобцов
        for (int j = 0; j < sizeSquareMatrix; j++)// здесь рассматриваем строки
        {
           columns[j] = squareMatrix[j][i];// закидываем в вектор столбец все элементы первого столбца где j номер строки 
            
        }
        if (i % 2 == 0)// если строка четная то сортируем в порядке возрастания
        {
            sort(columns.begin(), columns.end());
        }
        else// если не четная то в порядке убывания
        {
            sort(columns.rbegin(), columns.rend());
        }
        for (int k = 0; k < sizeSquareMatrix; k++)
        {
            squareMatrix[k][i] = columns[k];
        }
    }
    cout << endl;

    for (auto& i : squareMatrix)
    {
        for (auto& j : i)
        {
            cout << setw(4) << j;
        }
        cout << endl;
    }

    int sumRigth = 0;

    for (int i = 0; i < squareMatrix.size(); i++)// высчитываем в правый угол относительно диагонали
    {
        for (int j = i; j < sizeSquareMatrix; j++)
        {
            if (squareMatrix[i][j] > 0) sumRigth++;
        }

    }
    
    int sumLeft = 0;

    for (int i = sizeSquareMatrix; i > 0; i--)// высчитываем в левый угол относительно диагонали
    {
        for (int j = i; j > 0; j--)
        {
            if (squareMatrix[i-1][j-1] > 0) sumLeft++;
        }

    }
    cout << "Количество положительных элементов правой половины = " << sumRigth << "\t количество положительных элементов левой половины = " << sumLeft << endl;
    if (sumRigth > sumLeft) cout << "Правая половина содержит больше положительных элементов "<<endl;
    else cout << "Левая половина содержит больше положительных элементов" << endl;
}

int main()
{
    SetConsoleCP(1251);
    SetConsoleOutputCP(1251);

    task1();// доделать какието элементы матрицы
    cout << endl;
    task2();


}