f = open('/Users/frank/Downloads/mendian-store', 'r')
mst = f.readlines()
f.close()

f = open('/Users/frank/Downloads/mendian-shop', 'r')
msh = f.readlines()
f.close()

f = open('/Users/frank/Downloads/cloudqxt-com', 'r')
cc = f.readlines()
f.close()

f = open('/Users/frank/Downloads/danbangtuan-top', 'r')
dt = f.readlines()
f.close()

f = open('/Users/frank/Downloads/3306_all', 'r')
a = f.readlines()
f.close()

f = open('/Users/frank/Downloads/3307_all', 'r')
b = f.readlines()
f.close()

cloud_database = set()
for i in mst:
    cloud_database.add(i.strip())
for i in msh:
    cloud_database.add(i.strip())
for i in cc:
    cloud_database.add(i.strip())
for i in dt:
    cloud_database.add(i.strip())

source_database = set()
for i in a:
    source_database.add(i.strip())
for i in b:
    source_database.add(i.strip())

count = 0
for i in source_database:
    if not cloud_database.__contains__(i):
        count+=1
        print(i)

print(count)