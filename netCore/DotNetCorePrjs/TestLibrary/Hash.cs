namespace TestLibrary
{
    public class Hash
    {
        public static void Test()
        {
            string[] wordsDictionary = new string[50000];
            string word1 = "cat";
            string word2 = "abc";
            Insert(wordsDictionary, word1, word2);

            int findIndex = GenerateHashIndex(word1);

            int findIndex2 = GenerateHashIndex(word2);

        }

        private static void Insert(string[] wordsDictionary, string word1, string word2)
        {
            int word1Index = GenerateHashIndex(word1);
            int word2Index = GenerateHashIndex(word2);
            wordsDictionary[word1Index] = word1;
            wordsDictionary[word2Index] = word2;
        }

        private static int GenerateHashIndex(string word)
        {
            int hash = 0;

            foreach (char c in word)
            {
                hash = hash + GetIndex(c);
            }

            return hash;
        }
        private static int GetIndex(char letter)
        {
            switch (letter)
            {
                case 'a':
                    return 1;
                    
                    case 'b':   
                        return 2;
                    case 'c':
                        return 3;
                    case 'd':
                        return 4;
                    case 'e':
                        return 5;
                    case 'f':
                        return 6;
                case 'g':
                    return 7;
                    case 'h':
                        return 8;
                case 'i':
                    return 9;
                case 'j':
                    return 10;
                case 'k':
                    return 11;
                    case 'l':
                        return 12;
                case 'm':
                    return 13;
                    case 'n':
                        return 14;
                    case 'o':
                        return 15;
                case 'p':
                    return 16;
                case 'q':
                    return 17;
                    case 'r':
                        return 18;

                default:
                    break;
            }
            return 0;
        }
    }
}