char up = '(';
char down = ')';

int floor = 0;
int position = 1;

string directions = File.ReadAllText(Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "input.txt"));

foreach (char c in directions)
{
    if (c == up)
    {
        floor++;
    }
    else if (c == down)
    {
        floor--;
    }
    if (floor == -1)
    {
        Console.WriteLine($"Floor: {floor} Position: {position}");
        return;
    }
    position++;
}
