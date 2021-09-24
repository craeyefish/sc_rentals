from django.db import models


class Item(models.Model):
    name = models.CharField(max_length=200)
    create_date = models.DateTimeField('date published')
    value = models.IntegerField(default=0)

    def __str__(self):
        return self.name