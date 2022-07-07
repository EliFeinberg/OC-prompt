import json
import os
file_list = os.listdir('../source')
for files in file_list:
    if '.json' in files and files != "commands.json":
        f = open("../source/{}".format(files))
        data = json.load(f)
        print("\"{}\":".format(files[:len(files)-5])+ "{")
        for points in data:
            k = data[points]["Text"]
            v = data[points]["Description"]
            print("\tprompt.Suggest{"+"Text: \"{}\", Description: \"{}\"".format(k,v)+"},")
        print("},")
        f.close()
        

# f = open("../source/commands.json", "r")
# data = json.load(f)
# for points in data:
#     k = data[points]["Text"]
#     v = data[points]["Description"]
#     print("{"+"Text: \"{}\", Description: \"{}\"".format(k,v)+"},")
# close(f)