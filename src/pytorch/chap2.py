import os

from torchvision import models

response = dir(models)
# print(response)

alexNet = models.AlexNet()
# print(alexNet)

# 老版本写法
# resNet = models.resnet101(pretrained=True)
# 新版本写法
resNet = models.resnet101(weights = models.ResNet101_Weights.IMAGENET1K_V2)
# print(resNet)


# In[6]:
from torchvision import transforms
from PIL import Image
import torch
import os

preprocess = transforms.Compose([
        transforms.Resize(256),
        transforms.CenterCrop(224),
        transforms.ToTensor(),
        transforms.Normalize(
            mean=[0.485, 0.456, 0.406],
            std=[0.229, 0.224, 0.225]
        )])



with open('/Users/frank/Projects/dlwpt-code/data/p1ch2/imagenet_classes.txt') as f:
    labels = [line.strip() for line in f.readlines()]
path = "/Users/frank/Pictures/"
imgName = os.listdir(path)
for files in imgName:
    try:
        if not files.endswith(("jpg", "png")):
            continue;
        # img = Image.open("/Users/frank/Projects/dlwpt-code/data/p1ch2/bobby.jpg")
        img = Image.open(path + files)
        # img.show()
        img_t = preprocess(img)

        batch_t = torch.unsqueeze(img_t, 0)
        # 训练完train_datasets之后，model要来测试样本了。在model(test_datasets)之前，需要加上model.eval(). 否则的话，有输入数据，即使不训练，它也会改变权值。这是model中含有batch normalization层所带来的的性质。
        # eval（）时，pytorch会自动把BN和DropOut固定住，不会取平均，而是用训练好的值。不然的话，一旦test的batch_size过小，很容易就会被BN层导致生成图片颜色失真极大。eval（）在非训练的时候是需要加的，没有这句代码，一些网络层的值会发生变动，不会固定，你神经网络每一次生成的结果也是不固定的，生成质量可能好也可能不好。
        resNet.eval()
        output = resNet(batch_t)
        default, index = torch.max(output, 1)
        percentage = torch.nn.functional.softmax(output, dim=1)[0] * 100
        _, indices = torch.sort(output, descending=True)

        response = [(files, [(labels[idx], percentage[idx].item()) for idx in indices[0][:5]], labels[idx], percentage[idx].item()) for idx in indices[0][:5]]
        print(response)
    except:
        print(files, "exception~")
