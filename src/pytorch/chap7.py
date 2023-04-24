import torch
from torchvision import datasets

from torchvision import transforms

x = torch.tensor([1.0, 2.0, 3.0])
softmax = torch.nn.Softmax()
print(softmax(x))

tensor = torch.rand(2, 5, 5)
print(tensor.size())

tensor = tensor.squeeze(0)
print(tensor.size())

tensor = tensor.reshape(50)
print(tensor.size())


data_path = './'
print(dir(datasets))
cifar10 = datasets.CIFAR10(data_path, train = True, download = True)
cifar10_val = datasets.CIFAR10(data_path, train=False, download=True)

print(len(cifar10))
print(len(cifar10_val))

class_names = ['airplane','automobile','bird','cat','deer',
               'dog','frog','horse','ship','truck']

img, label = cifar10[1000]
print(label, class_names[label])

print(dir(transforms))

print("nn", dir(torch.nn))

tensor_cifar10 = datasets.CIFAR10(data_path, train=True, download=False,
                              transform=transforms.ToTensor())

model = torch.nn.Sequential(
    torch.nn.Linear(3072, 512),
    torch.nn.Tanh(),
    torch.nn.Linear(512, 2),
    torch.nn.Softmax(dim = 1)
)
