import datetime

from django.db import models
from django.utils import timezone


class Item(models.Model):
    name = models.CharField(max_length=200)
    create_date = models.DateTimeField('date published')
    value = models.IntegerField(default=0)

    def __str__(self):
        return self.name

    def created_recently(self):
        if timezone.now() - self.create_date > datetime.timedelta(days=1):
            return False
        return True