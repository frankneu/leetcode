import torch
from torchvision import datasets
from torchvision import transforms
from matplotlib import pyplot as plt

data_path = './'
cifar10 = datasets.CIFAR10(data_path, train = True, download = True,
                           transform=transforms.ToTensor())
cifar10_val = datasets.CIFAR10(data_path, train = False, download = True,
                               transform=transforms.ToTensor())

class_names = ['airplane','automobile','bird','cat','deer',
               'dog','frog','horse','ship','truck']
tensor, label = cifar10[99]
print("source", tensor.size(), class_names[label])

convolution = torch.nn.Conv2d(3, 16, kernel_size=3, padding = 1)
pool = torch.nn.MaxPool2d(2)
tanh = torch.nn.Tanh()
# 处理整个网络
out = convolution(tensor)
print("convolution", out.size())
out = tanh(out)
print("tanh", out.size())
out = pool(out)
print("pooling", out.size())
convolution = torch.nn.Conv2d(16, 8, kernel_size=3, padding = 1)
out = convolution(out)
print("convolution", out.size())
out = tanh(out)
print("tanh", out.size())
out = pool(out)
print("pooling", out.size())
out = out.view(-1, 8 * 8 * 8)
print("view", out.size())
linear = torch.nn.Linear(8 * 8 * 8, 32)
out = linear(out)
print("linear", out.size())
out = tanh(out)
print("tanh", out.size())
linear = torch.nn.Linear(32, 2)
out = linear(out)
print("linear", out.size())

