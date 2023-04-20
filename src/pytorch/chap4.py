# import imageio
# import torch
# from torchvision.transforms import InterpolationMode
# import os
# import torchvision.transforms as transforms

# img_array = imageio.v2.imread("/Users/frank/Pictures/WanderEarth2.jpg")
# print(img_array.shape)
#
# img_tensor = torch.from_numpy(img_array)
# print(img_tensor.shape)
# img_tensor_permute = img_tensor.permute(2,1,0)
# print(img_tensor_permute.shape)
#
#
#
# picture_dir = "/Users/frank/Pictures"
# imgName = os.listdir(picture_dir)
# batch = torch.zeros(imgName.__sizeof__(), 3, 3000, 2000, dtype=torch.uint8)
# transform = transforms.Compose([
#     transforms.Resize((3000,2000), interpolation=InterpolationMode.BICUBIC) # 统一图片大小为224 x 224，并保持比例不变
#     # other transformations...
# ])
#
# for i, files in enumerate(imgName):
#     if not files.endswith(("jpg", "png")):
#         continue;
#     image_array = imageio.v2.imread(os.path.join(picture_dir, files))
#     img_temporary = torch.from_numpy(img_array)
#     img_temporary = img_temporary.permute(2,1,0)
#     img_temporary = transform(img_temporary)
#     # 图片只获取前三个channel，因为部分图片有alpha通道，这是一个透明度通道
#     img_temporary = img_temporary[:3]
#     batch[i] = img_temporary
# print("before", batch[0][0][0][0])
# batch = batch.float()
# print("float", batch[0][0][0][0])
# batch /= 10.0
# print("div ten", batch[0][0][0][0])


# array = [1,2,3,4,5,6,7,8,9]
# print(array)
# print(array[-1])
# print(array[-1:])
# print(array[:-1])

import csv
import numpy as np
import torch

# 评估酒水质量
wine_path = "/Users/frank/Projects/dlwpt-code/data/p1ch4/tabular-wine/winequality-white.csv"
wineq_numpy = np.loadtxt(wine_path, dtype=np.float32, delimiter=";",skiprows=1)
col_list = next(csv.reader(open(wine_path), delimiter=';'))
wineq = torch.from_numpy(wineq_numpy)
print(wineq.size())

data = wineq[:, :-1] # 参数
target = wineq[:, -1] # 目标值
target = wineq[:, -1].long()
# 改为one-hot方式编码，更适合质量这种不均衡的评估方式，因为这些结果没有严格的倍数关系
target_onehot = torch.zeros(target.shape[0], 10)
target_onehot.scatter_(1, target.unsqueeze(1), 1.0)
target_unsqueezed = target.unsqueeze(1) # 增加一个维度，方便处理
data_mean = torch.mean(data, dim = 0)
data_var = torch.var(data, dim=0)
data_normalized = (data - data_mean) / torch.sqrt(data_var)
print(data_normalized)

bad_data = data[target <= 3]
mid_data = data[(target > 3) & (target < 7)]
good_data = data[target >= 7]

bad_mean = torch.mean(bad_data, dim=0)
mid_mean = torch.mean(mid_data, dim=0)
good_mean = torch.mean(good_data, dim=0)

for i, args in enumerate(zip(col_list, bad_mean, mid_mean, good_mean)):
    print('{:2} {:20} {:6.2f} {:6.2f} {:6.2f}'.format(i, *args))

