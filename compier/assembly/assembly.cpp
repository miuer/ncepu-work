#include <cstdio>
#include <vector>
#include <cstdlib>
#include <algorithm>
#include <string>
#include <stack>
#include <iostream>
#include <utility>
#define In() freopen("in.txt", "r", stdin)
#define Out() freopen("assembly.txt", "w", stdout)

using namespace std;

/*

E->TG
G->+TG|-TG|ε
T->FS
S->*FS|/FS|ε
F->(E)|i

文法调整

E -> T A
A -> + T GEN(+) A
A -> - T GEN(-) A
A -> ε
T -> FB
B -> * F GEN(*) B
B -> / F GEN(/) B
B -> ε
F -> I PUSH(I)
F -> (E)

*/


struct Qt //定义四元式结构体
{
	string op, a, b, c;
};

string input_str;
int i = 0, num = 1;
string ch;
bool flag = 0; // flag = 1 表示当前算数表达式串无法识别
vector<pair<string, char> > str;
vector<Qt> ans;
stack<string> s;

void To_String() //将输入字符串转换为TOKEN串
{
	int i = 0;
	while (i < (int)input_str.size())
	{
		if (input_str[i] == ' ') // 空字符跳过，不进行处理
			i++;
		else if ((input_str[i] >= 'a' && input_str[i] <= 'z') || (input_str[i] >= 'A' && input_str[i] <= 'Z') || (input_str[i] >= '0' && input_str[i] <= '9')) //当前是字符或者数字
		{
			string tmp;
			for (; i < (int)input_str.size(); i++)
			{
				if (!((input_str[i] >= 'a' && input_str[i] <= 'z') || (input_str[i] >= 'A' && input_str[i] <= 'Z') || (input_str[i] >= '0' && input_str[i] <= '9') || (input_str[i] == '.')))
					break;
				else
					tmp.push_back(input_str[i]);
			}
			str.push_back(make_pair(tmp, 'I'));
		}
		else //其他情况（单运算符），直接将字符放入
		{
			string tmp;
			tmp.push_back(input_str[i]);
			str.push_back(make_pair(tmp, input_str[i]));
			i++;
		}
	}
}

void target(Qt qt)
{
	if (qt.op == "+")
	{
		cout << "mov eax," << qt.a << endl;
		cout << "add eax," << qt.b << endl;
		cout << "mov " << qt.c << ", eax" << endl;
	}
	else if (qt.op == "-")
	{
		cout << "mov eax," << qt.a << endl;
		cout << "sub eax," << qt.b << endl;
		cout << "mov " << qt.c << ", eax" << endl;
	}
	else if (qt.op == "*")
	{
		cout << "mov eax," << qt.a << endl;
		cout << "mul eax," << qt.b << endl;
		cout << "mov " << qt.c << ", eax" << endl;
	}
	else if (qt.op == "/")
	{
		cout << "mov eax," << qt.a << endl;
		cout << "div eax," << qt.b << endl;
		cout << "mov " << qt.c << ", eax" << endl;
	}
}

int main() //LL(1)方法
{
	In();
	Out();
	input_str.clear();

	while (getline(cin, input_str))
	{
		//stack<string>s;
		stack<string> sem;
		ans.clear();
		str.clear();
		i = 0;
		flag = 0;
		num = 1;
		cout << input_str << endl;
		To_String();
		//将输入的串转换为TOKEN串
		
		s.push("#");
		s.push("E"); //初始化栈

		i = 0;
		while (!s.empty() && i < (int)str.size())
		{
			//print_stack();
			ch = s.top();

			string tmp1;
			tmp1.push_back(str[i].second);

			// 迭代出口
			if (ch == "#")
			{
				// 输入串含有不合法字符 '#'，
				if (str[i].second == '#')
				{
					s.pop();
					flag = 1;
					break;
				}
				else
					break;
			}
			else if (ch == tmp1)
			{
				s.pop();
				i++;
			}
			else if (ch == "E") //当前栈顶元素为E
			{
				if (str[i].second == 'I' || str[i].second == '(')
				{
					s.pop();
					s.push("A");
					s.push("T");
				} //将 1 号产生式逆序压栈
				else
					break;
			}
			else if (ch == "A") //当前栈顶元素为A
			{
				if (str[i].second == '+')
				{
					s.pop();
					s.push("A");
					s.push("GEN(+)");
					s.push("T");
					s.push("+");
				} //将 2 号产生式逆序压栈
				else if (str[i].second == '-')
				{
					s.pop();
					s.push("A");
					s.push("GEN(-)");
					s.push("T");
					s.push("-");
				} //将 3 号产生式逆序压栈
				else if (str[i].second == ')' || str[i].second == '#')
				{
					s.pop();
				} //将 4 号产生式逆序压栈，空串只进行弹栈
				else
					break;
			}
			else if (ch == "T")
			{
				if (str[i].second == 'I' || str[i].second == '(')
				{
					s.pop();
					s.push("B");
					s.push("F");
				} //将 5 号产生式逆序压栈
				else
					break;
			}
			else if (ch == "B")
			{
				if (str[i].second == '*')
				{
					s.pop();
					s.push("B");
					s.push("GEN(*)");
					s.push("F");
					s.push("*");
				} //将 6 号产生式逆序压栈
				else if (str[i].second == '/')
				{
					s.pop();
					s.push("B");
					s.push("GEN(/)");
					s.push("F");
					s.push("/");
				} //将 7 号产生式逆序压栈
				else if (str[i].second == '+' || str[i].second == '-' || str[i].second == ')' || str[i].second == '#')
				{
					s.pop();
				} //将 8 号产生式逆序压栈，空串只进行弹栈
				else
					break;
			}
			else if (ch == "F")
			{
				if (str[i].second == 'I')
				{
					string tmp = "PUSH-" + str[i].first;
					s.pop();
					s.push(tmp);
					s.push("I");
				} //将 9 号产生式逆序压栈
				else if (str[i].second == '(')
				{
					s.pop();
					s.push(")");
					s.push("E");
					s.push("(");
				} //将 10 号产生式逆序压栈
				else
					break;
			}
			else if (ch[0] == 'P') //PUSH操作
			{
				string tmp;
				for (int k = 5; k < (int)ch.size(); k++)
				{
					tmp.push_back(ch[k]);
				}
				sem.push(tmp);
				s.pop();
			}
			else if (ch[0] == 'G') //生成四元式操作
			{
				string tmp;

				// Ascll 1 --- 49
				tmp.push_back('t');
				tmp.push_back(48 + num);
				//cout << tmp <<endl;
				string res1 = sem.top();
				sem.pop();
				string res2 = sem.top();
				sem.pop();

				Qt qt;
				qt.op = ch[4];
				qt.a = res2;
				qt.b = res1;
				qt.c = tmp;
				ans.push_back(qt);

				sem.push(tmp);
				num++;
				s.pop();
			}
		}

		if (flag == 1)
		{
			printf("Yes\n");
			printf("\n");

			for (int k = 0; k < (int)ans.size(); k++)
			{
				cout << '(';
				cout << ans[k].op;
				cout << ',';
				cout << ans[k].a;
				cout << ',';
				cout << ans[k].b;
				cout << ',';
				cout << ans[k].c;
				cout << ')' << endl;
			}
			
			printf("\n");

			for (int k = 0; k < (int)ans.size(); k++)
			{
				target(ans[k]);
			}
		}
		else
			printf("No\n");
	}

	return 0;
}